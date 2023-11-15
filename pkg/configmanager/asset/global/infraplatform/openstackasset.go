/*
Copyright 2023 KylinSoft  Co., Ltd.

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

package infraplatform

var OpenStackConfig *OpenStackAsset

// ========== 包方法 ==========

func GetOpenStackConfig() (*OpenStackAsset, error) {
	return OpenStackConfig, nil
}

// ========== 模块方法 ==========

type OpenStackAsset struct {
	Auth_Url string
}

// TODO: Init inits the openstack asset.
func (opa *OpenStackAsset) Initial() error {
	opa.Auth_Url = "http://"

	OpenStackConfig = opa
	return nil
}

// TODO: Delete deletes the openstack asset.
func (opa *OpenStackAsset) Delete() error {
	return nil
}

// TODO: Persist persists the openstack asset.
func (opa *OpenStackAsset) Persist() error {
	return nil
}
