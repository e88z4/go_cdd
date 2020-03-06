/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EmailNotificationServerDto struct {
	ConnectivityArguments []ArgumentDto `json:"connectivityArguments,omitempty"`
	DestinationEmail string `json:"destinationEmail,omitempty"`
	Host string `json:"host,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsEnabled bool `json:"isEnabled,omitempty"`
	IsSSL bool `json:"isSSL,omitempty"`
	LastConnectivityStatus string `json:"lastConnectivityStatus,omitempty"`
	LastConnectivityTestDate int64 `json:"lastConnectivityTestDate,omitempty"`
	LastConnectivityTestMessage string `json:"lastConnectivityTestMessage,omitempty"`
	Password string `json:"password,omitempty"`
	Port int32 `json:"port,omitempty"`
	Sender string `json:"sender,omitempty"`
	SenderDisplayName string `json:"senderDisplayName,omitempty"`
	User string `json:"user,omitempty"`
}
