/*
 * Copyright 2020 The Compass Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"reflect"

	"github.com/kyma-incubator/compass/components/admiral-watcher/pkg/env"
	"github.com/kyma-incubator/compass/components/admiral-watcher/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

type Validatable interface {
	Validate() error
}

type Config struct {
	Log *log.Config `mapstructure:"log"`
}

func AddPFlags(set *pflag.FlagSet) {
	env.CreatePFlags(set, DefaultConfig())
	env.CreatePFlagsForConfigFile(set)
}

func DefaultConfig() *Config {
	return &Config{
		Log: log.DefaultConfig(),
	}
}

func New(env env.Environment) (*Config, error) {
	config := DefaultConfig()
	if err := env.Unmarshal(config); err != nil {
		return nil, errors.Wrapf(err, "error loading cfg")
	}

	return config, nil
}

func (c *Config) Validate() error {
	validatableFields := make([]Validatable, 0, 0)

	v := reflect.ValueOf(*c)
	for i := 0; i < v.NumField(); i++ {
		field, ok := v.Field(i).Interface().(Validatable)
		if ok {
			validatableFields = append(validatableFields, field)
		}
	}

	for _, item := range validatableFields {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
