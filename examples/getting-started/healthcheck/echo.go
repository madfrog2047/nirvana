/*
Copyright 2018 Caicloud Authors

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

package main

import (
	"context"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/config"
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/errors"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/plugins/healthcheck"
)

var echo = definition.Descriptor{
	Path:        "/echo",
	Description: "Echo API",
	Definitions: []definition.Definition{
		{
			Method:   definition.Get,
			Function: Echo,
			Consumes: []string{definition.MIMEAll},
			Produces: []string{definition.MIMEJSON},
			Parameters: []definition.Parameter{
				{
					Source:      definition.Query,
					Name:        "msg",
					Description: "Corresponding to the second parameter",
				},
			},
			Results: []definition.Result{
				{
					Destination: definition.Data,
					Description: "Corresponding to the first result",
				},
				{
					Destination: definition.Error,
					Description: "Corresponding to the second result",
				},
			},
		},
	},
}

// API function.
func Echo(ctx context.Context, msg string) (string, error) {
	return msg, nil
}

func main() {
	cmd := config.NewDefaultNirvanaCommand()
	cfg := nirvana.NewDefaultConfig()
	cfg.Configure(
		nirvana.Descriptor(echo),
		healthcheck.CheckerWithType(func(ctx context.Context, checkType string) error {
			switch checkType {
			case healthcheck.LivenessCheck:
				// do something
				return nil
			case healthcheck.ReadinessCheck:
				// do something
				return nil
			default:
				return errors.BadRequest.Build("error", "unknown type ${type}").Error(checkType)
			}
		}),
		// healthcheck.Checker(func(ctx context.Context) error {
		// 	return nil
		// }),
	)
	if err := cmd.ExecuteWithConfig(cfg); err != nil {
		log.Fatal(err)
	}
}
