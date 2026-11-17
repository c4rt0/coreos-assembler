// Copyright 2026 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kola

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/coreos/coreos-assembler/mantle/kola/register"
)

func TestRegisterExternalTestCPUCountMetadata(t *testing.T) {
	tests := []struct {
		name      string
		script    string
		kolaJSON  string
		wantCPU   int
		platforms []string
	}{
		{
			name:      "kola.json",
			script:    "#!/bin/bash\necho test\n",
			kolaJSON:  `{"platforms":"qemu","exclusive":true,"cpuCount":2}`,
			wantCPU:   2,
			platforms: []string{"qemu"},
		},
		{
			name:      "inline YAML",
			script:    "#!/bin/bash\n## kola:\n##   platforms: qemu\n##   cpuCount: 2\necho test\n",
			wantCPU:   2,
			platforms: []string{"qemu"},
		},
		{
			name:      "inline JSON",
			script:    "#!/bin/bash\n# kola: {\"platforms\":\"qemu\",\"cpuCount\":2}\necho test\n",
			wantCPU:   2,
			platforms: []string{"qemu"},
		},
		{
			name:   "no metadata",
			script: "#!/bin/bash\necho test\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			if tt.kolaJSON != "" {
				if err := os.WriteFile(filepath.Join(dir, "kola.json"), []byte(tt.kolaJSON), 0644); err != nil {
					t.Fatal(err)
				}
			}
			if err := os.WriteFile(filepath.Join(dir, "test.sh"), []byte(tt.script), 0755); err != nil {
				t.Fatal(err)
			}

			children, err := os.ReadDir(dir)
			if err != nil {
				t.Fatal(err)
			}
			name := "test-cpucount-" + tt.name
			t.Cleanup(func() { delete(register.Tests, name) })
			if err := registerTestDir(dir, name, children); err != nil {
				t.Fatal(err)
			}
			test := register.Tests[name]
			if test == nil {
				t.Fatal("test was not registered")
			}
			if test.MachineOptions.CPUCount != tt.wantCPU {
				t.Errorf("CPUCount = %d, want %d", test.MachineOptions.CPUCount, tt.wantCPU)
			}
			if !slices.Equal(test.Platforms, tt.platforms) {
				t.Errorf("Platforms = %v, want %v", test.Platforms, tt.platforms)
			}
			if test.NonExclusive {
				t.Error("test is non-exclusive, want exclusive")
			}
		})
	}
}
