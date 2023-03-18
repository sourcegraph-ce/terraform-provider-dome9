package dome9

import (
	log "github.com/sourcegraph-ce/logrus"

	"github.com/dome9/dome9-sdk-go/dome9"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/aws"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/azure"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/gcp"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupaws"
	"github.com/dome9/dome9-sdk-go/services/cloudsecuritygroup/securitygroupazure"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_notification"
	"github.com/dome9/dome9-sdk-go/services/compliance/continuous_compliance_policy"
	"github.com/dome9/dome9-sdk-go/services/iplist"
	"github.com/dome9/dome9-sdk-go/services/organizationalunits"
	"github.com/dome9/dome9-sdk-go/services/roles"
	"github.com/dome9/dome9-sdk-go/services/rulebundles"
	"github.com/dome9/dome9-sdk-go/services/users"
)

func init() {
	// remove timestamp from Dome9 provider logger, use the timestamp from the default terraform logger
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

type Client struct {
	iplist                           iplist.Service
	cloudaccountAWS                  aws.Service
	cloudaccountAzure                azure.Service
	cloudaccountGCP                  gcp.Service
	cloudaccountKubernetes           k8s.Service
	continuousCompliancePolicy       continuous_compliance_policy.Service
	continuousComplianceNotification continuous_compliance_notification.Service
	ruleSet                          rulebundles.Service
	awsSecurityGroup                 securitygroupaws.Service
	role                             roles.Service
	organizationalUnit               organizationalunits.Service
	azureSecurityGroup               securitygroupazure.Service
	users                            users.Service
}

type Config struct {
	AccessID  string
	SecretKey string
	BaseURL   string
}

func (c *Config) Client() (*Client, error) {
	config, err := dome9.NewConfig(c.AccessID, c.SecretKey, c.BaseURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		iplist:                           *iplist.New(config),
		cloudaccountAWS:                  *aws.New(config),
		cloudaccountAzure:                *azure.New(config),
		cloudaccountGCP:                  *gcp.New(config),
		cloudaccountKubernetes:           *k8s.New(config),
		continuousCompliancePolicy:       *continuous_compliance_policy.New(config),
		continuousComplianceNotification: *continuous_compliance_notification.New(config),
		ruleSet:                          *rulebundles.New(config),
		awsSecurityGroup:                 *securitygroupaws.New(config),
		role:                             *roles.New(config),
		organizationalUnit:               *organizationalunits.New(config),
		azureSecurityGroup:               *securitygroupazure.New(config),
		users:                            *users.New(config),
	}

	log.Println("[INFO] initialized Dome9 client")
	return client, nil
}
