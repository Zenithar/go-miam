// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

// ApplicationCreated is raised when an application domain has been created.
func ApplicationCreated(id, label string) interface{} {
	return nil
}

// ApplicationActivated is raised when an application domain has been activated.
func ApplicationActivated(id string) interface{} {
	return nil
}

// ApplicationDisabled is raised when an application domain has been disabled.
func ApplicationDisabled(id string) interface{} {
	return nil
}

// ApplicationLabelChanged is raised when an application label attribute has been changed.
func ApplicationLabelChanged(id, old, new string) interface{} {
	return nil
}
