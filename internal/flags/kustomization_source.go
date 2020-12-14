/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flags

import (
	"fmt"
	"strings"

	"github.com/fluxcd/flux2/internal/utils"
	sourcev1 "github.com/fluxcd/source-controller/api/v1beta1"
)

var supportedKustomizationSourceKinds = []string{sourcev1.GitRepositoryKind, sourcev1.BucketKind}

type KustomizationSource struct {
	Kind string
	Name string
}

func (s *KustomizationSource) String() string {
	if s.Name == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", s.Kind, s.Name)
}

func (s *KustomizationSource) Set(str string) error {
	if strings.TrimSpace(str) == "" {
		return fmt.Errorf("no Kustomization source given, please specify %s",
			s.Description())
	}

	sourceKind, sourceName := utils.ParseObjectKindName(str)
	if sourceName == "" {
		return fmt.Errorf("no name given for source of kind '%s'", sourceKind)
	}
	if sourceKind == "" {
		if utils.ContainsItemString(supportedKustomizationSourceKinds, sourceName) {
			return fmt.Errorf("no name given for source of kind '%s'", sourceName)
		}
		sourceKind = sourcev1.GitRepositoryKind
	}
	if !utils.ContainsItemString(supportedKustomizationSourceKinds, sourceKind) {
		return fmt.Errorf("source kind '%s' is not supported, must be one of: %s",
			sourceKind, strings.Join(supportedKustomizationSourceKinds, ", "))
	}

	s.Name = sourceName
	s.Kind = sourceKind

	return nil
}

func (s *KustomizationSource) Type() string {
	return "kustomizationSource"
}

func (s *KustomizationSource) Description() string {
	return fmt.Sprintf(
		"source that contains the Kubernetes manifests in the format '[<kind>/]<name>', "+
			"where kind must be one of: (%s), if kind is not specified it defaults to GitRepository",
		strings.Join(supportedKustomizationSourceKinds, ", "),
	)
}
