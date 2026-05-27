// Copyright (c) Citrix, Inc.

package common

import (
	"context"
	"fmt"
	"log"
	"time"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type WaitIpConfig struct {
	// Amount of time to wait for VM's IP, similar to 'ssh_timeout'.
	// Defaults to `30m` (30 minutes). Refer to the Golang
	// [ParseDuration](https://golang.org/pkg/time/#ParseDuration)
	// documentation for full details.
	WaitTimeout int `mapstructure:"ip_wait_timeout"`

	OperationType elmsoap.ApplayeringOperationType
	LayerName     string
}

type StepWaitForIp struct {
	Config *WaitIpConfig
}

func (c *WaitIpConfig) Prepare() []error {
	var errs []error
	if c.WaitTimeout == 0 {
		c.WaitTimeout = 30
	}

	return errs
}

func (s *StepWaitForIp) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	helper := state.Get("soap_helper").(*elmsoap.SoapHelper)
	workTicketId := state.Get("WORK_TICKET_ID").(int64)
	UiSayf(ui, "Waiting for IP for WorkTicketId: %d", workTicketId)
	generated_data := state.Get("generated_data").(map[string]any)
	generated_data["WORK_TICKET_ID"] = workTicketId

	var ip string
	var err error

	sub, cancel := context.WithCancel(ctx)
	waitDone := make(chan bool, 1)
	defer func() {
		cancel()
	}()

	go func() {
		ip, err = doGetIp(ui, sub, helper, workTicketId)
		waitDone <- true
	}()

	UiSayf(ui, "[INFO] Waiting for IP, up to total timeout: %d minutes.", s.Config.WaitTimeout)
	timeout := time.After(time.Duration(s.Config.WaitTimeout) * time.Minute)
	for {
		select {
		case <-timeout:
			cancel()
			<-waitDone
			if ip != "" {
				state.Put("ip", ip)
				UiSayf(ui, "[WARN] API timeout waiting for IP but one IP was found. Using IP: %s", ip)
				return multistep.ActionContinue
			}
			err := fmt.Errorf("timeout waiting for IP address")
			state.Put("error", err)
			ui.Errorf("%s", err)
			return multistep.ActionHalt
		case <-ctx.Done():
			cancel()
			UiSay(ui, "[WARN] Interrupt detected, quitting waiting for IP.")
			return multistep.ActionHalt
		case <-waitDone:
			if err != nil {
				state.Put("error", err)
				return multistep.ActionHalt
			}
			state.Put("ip", ip)
			UiSayf(ui, "IP address: %v", ip)
			return multistep.ActionContinue
		case <-time.After(1 * time.Second):
			if _, ok := state.GetOk(multistep.StateCancelled); ok {
				return multistep.ActionHalt
			}
		}
	}
}

func doGetIp(ui packersdk.Ui, ctx context.Context, helper *elmsoap.SoapHelper, workTicketId int64) (string, error) {
	interval := 10 * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("cancelled waiting for IP address")
		case <-ticker.C:
			// Call GetWorkTicketId periodically
			ip, err := helper.GetIpByWorkTicketId(workTicketId)
			if err != nil {
				// Try to find the work ticket ID from the complete task info.
				workTicketResult, err := helper.GetTaskCompletedFilter(workTicketId)
				if err != nil {
					log.Printf("[ERROR] Failed to get completed task info for work ticket ID %d: %v", workTicketId, err)
					return "", fmt.Errorf("[ERROR] Failed to get IP address for work ticket ID %d: %v", workTicketId, err)
				}
				if (*workTicketResult.State) != "Success" {
					log.Printf("[ERROR] Work ticket ID %d did not complete successfully. State: %s. Ticket: %v", workTicketId, *workTicketResult.State, workTicketResult)
					workItemStatus := ""
					for _, workItem := range workTicketResult.WorkItems.WorkItemResult {
						if *workItem.ItemType != "WorkItem" {
							workItemStatus = workItem.Status
							break
						}
					}
					return "", fmt.Errorf("[ERROR] Work ticket ID %d did not complete successfully. WorkTicketResultState: %s\n WorkItemStatus:%s", workTicketId, *workTicketResult.State, workItemStatus)
				}
				return "", fmt.Errorf("[ERROR] Failed to get work ticket ID %d: %v", workTicketId, err)
			}
			// Return the IP if it is valid
			if ip != "" {
				UiSayf(ui, "Found IP address: %s for work ticket id: %d", ip, workTicketId)
				return ip, nil
			}
			log.Printf("IP Address is not ready yet, retrying...")
		}
	}
}

func (s *StepWaitForIp) Cleanup(state multistep.StateBag) {}
