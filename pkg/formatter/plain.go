package formatter

import (
	"fmt"

	"github.com/rikatz/kubepug/pkg/kubepug"
)

type plain struct{}

func newPlainFormatter() Formatter {
	return &plain{}
}

func (f *plain) Output(results kubepug.Result) ([]byte, error) {
	s := fmt.Sprintf("RESULTS:\nDeprecated APIs:\n\n")
	for _, api := range results.DeprecatedAPIs {
		s = fmt.Sprintf("%s%s found in %s/%s\n", s, api.Kind, api.Group, api.Version)
		if api.Description != "" {
			s = fmt.Sprintf("%sDescription: %s\n", s, api.Description)
		}
		items := listItems(api.Items)
		s = fmt.Sprintf("%s%s\n", s, items)
	}
	s = fmt.Sprintf("%s\nDeleted APIs:\n\n", s)
	for _, api := range results.DeletedAPIs {
		s = fmt.Sprintf("%s%s found in %s/%s\n", s, api.Kind, api.Group, api.Version)
		items := listItems(api.Items)
		s = fmt.Sprintf("%s%s\n", s, items)
	}
	return []byte(s), nil
}

func listItems(items []kubepug.DeprecatedItem) string {
	s := fmt.Sprintf("")
	for _, i := range items {
		if i.Namespace != "" {
			s = fmt.Sprintf("%s%s: %s namespace: %s\n", s, i.Kind, i.Name, i.Namespace)
		} else {
			s = fmt.Sprintf("%s%s: %s \n", s, i.Kind, i.Name)
		}
	}
	return s
}