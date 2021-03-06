/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
// Health The operational health of an entity 
type Health string

// List of Health
const (
	HEALTH_ONLINE Health = "online"
	HEALTH_OFFLINE Health = "offline"
	HEALTH_UNKNOWN Health = "unknown"
)
