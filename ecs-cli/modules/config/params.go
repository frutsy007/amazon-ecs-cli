// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package config

import (
	"github.com/Sirupsen/logrus"
	ecscli "github.com/aws/amazon-ecs-cli/ecs-cli/modules/commands"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/urfave/cli"
)

// CliParams saves config to create an aws service clients
type CliParams struct {
	Cluster string
	Session *session.Session
}

// NewCliParams creates a new ECSParams object from the config file.
func NewCliParams(context *cli.Context, rdwr ReadWriter) (*CliParams, error) {
	ecsConfig, err := rdwr.GetConfig()
	if err != nil {
		logrus.Error("Error loading config: ", err)
		return nil, err
	}

	// The global --region flag has the highest precedence to set ecs-cli region config.
	regionFromFlag := context.GlobalString(ecscli.RegionFlag)
	if regionFromFlag != "" {
		ecsConfig.Region = regionFromFlag
	}

	svcSession, err := ecsConfig.ToAWSSession()
	if err != nil {
		return nil, err
	}

	return &CliParams{
		Cluster: ecsConfig.Cluster,
		Session: svcSession,
	}, nil
}
