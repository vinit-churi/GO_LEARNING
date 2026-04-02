//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"strings"
)

type AuditInfo struct {
	Owner string
	Team  string
}

func (a AuditInfo) FullName() string {
	return strings.TrimSpace(a.Owner) + " @ " + strings.TrimSpace(a.Team)
}

type DeployService struct {
	AuditInfo
	Name string
}

func (d DeployService) Summary() string {
	return fmt.Sprintf("%s owned by %s", d.Name, d.FullName())
}

func (d *DeployService) Touch(name string) {
	d.Owner = strings.TrimSpace(name)
}

func main() {
	service := DeployService{AuditInfo: AuditInfo{Owner: "Ria", Team: "platform"}, Name: "billing-api"}
	service.Touch(" Riya ")
	fmt.Println(service.Summary())
}
