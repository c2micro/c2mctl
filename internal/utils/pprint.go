package utils

import (
	"fmt"
	"strings"
	"time"

	managementv1 "github.com/c2micro/c2mshr/proto/gen/management/v1"
	"github.com/docker/go-units"
)

func PrintOperator(o *managementv1.Operator) string {
	var s strings.Builder
	// username
	s.WriteString(fmt.Sprintf("Username: %s\n", o.Username))
	// token
	if o.Token != nil && o.Token.GetValue() != "" {
		s.WriteString(fmt.Sprintf("Token:    %s\n", o.Token.GetValue()))
	} else {
		s.WriteString("Token:    <none>\n")
	}
	// last
	if !o.Last.AsTime().IsZero() {
		s.WriteString(fmt.Sprintf("Last:     %s ago", units.HumanDuration(time.Since(o.Last.AsTime()))))
	} else {
		s.WriteString("Last:     <never>")
	}
	return s.String()
}

func PrintListener(l *managementv1.Listener) string {
	var s strings.Builder
	// id
	s.WriteString(fmt.Sprintf("ID:    %d\n", l.Lid))
	// token
	if l.Token != nil && l.Token.GetValue() != "" {
		s.WriteString(fmt.Sprintf("Token: %s\n", l.Token.GetValue()))
	} else {
		s.WriteString("Token: <none>\n")
	}
	// name
	if l.Name != nil {
		s.WriteString(fmt.Sprintf("Name:  %s\n", l.Name.GetValue()))
	} else {
		s.WriteString("Name:  <none>\n")
	}
	// ip
	if l.Ip != nil {
		s.WriteString(fmt.Sprintf("IP:    %s\n", l.Ip.GetValue()))
	} else {
		s.WriteString("IP:    <none>\n")
	}
	// port
	if l.Port != nil {
		s.WriteString(fmt.Sprintf("Port:  %d\n", l.Port.GetValue()))
	} else {
		s.WriteString("Port:  <none>\n")
	}
	// last
	if !l.Last.AsTime().IsZero() {
		s.WriteString(fmt.Sprintf("Last:  %s ago", units.HumanDuration(time.Since(l.Last.AsTime()))))
	} else {
		s.WriteString("Last:  <never>")
	}

	return s.String()
}
