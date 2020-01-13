package resources

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-grep/pkg/constants"
	"github.com/guessi/kubectl-grep/pkg/options"
	"github.com/guessi/kubectl-grep/pkg/utils"
)

// Secrets - a public function for searching secrets with keyword
func Secrets(opt *options.SearchOptions, keyword string) {
	var secretInfo string

	secretList := utils.SecretList(opt)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.SecretHeader)

	for _, s := range secretList.Items {
		// return all secrets under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if !match {
				continue
			}
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		secretInfo = fmt.Sprintf(constants.SecretRowTemplate,
			s.Namespace,
			s.Name,
			s.Type,
			len(s.Data),
			age,
		)

		fmt.Fprintln(w, secretInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}