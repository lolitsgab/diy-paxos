package utils

import (
	"fmt"
	"github.com/avast/retry-go"
	"log"
	"net"
	"strings"
	"time"
)

func getIpsWithoutPorts(ips string) []string {
	// Split the string into individual IP addresses
	ipList := strings.Split(ips, ",")

	// Create a slice to hold the IP addresses
	parsedIPs := make([]string, len(ipList))

	// Loop through each IP address in the list and add it to the parsedIPs slice
	for i, addrStr := range ipList {
		parsedIPs[i] = strings.TrimSpace(addrStr)
	}

	return parsedIPs
}

// GetReplicaIPs resolves the IP addressed of all replicas.
func DiscoverReplicas(discoveryHost, host string, retries uint, backOffDelay time.Duration) ([]string, error) {
	var reps []string
	err := retry.Do(func() error {
		resp, err := net.LookupIP(discoveryHost)
		if err != nil {
			log.Printf("%v could not get replicas from %v: %v", host, discoveryHost, err)
			return err
		}
		for _, ip := range resp {
			reps = append(reps, ip.String()+":"+"8080")
		}
		log.Printf("%v has replicas: %v", host, reps)
		return nil
	},
		retry.Attempts(retries),
		retry.Delay(backOffDelay),
		retry.DelayType(retry.BackOffDelay),
	)
	return reps, err
}

func ParseReplicaString(replicas string) ([]string, error) {
	ipList := strings.Split(replicas, ",")

	// Create a slice to hold the parsed IP addresses
	replicasSlice := make([]string, len(ipList))

	// Loop through each IP address in the list and parse it
	for i, ipStr := range ipList {
		ipAddr, portStr, err := net.SplitHostPort(strings.TrimSpace(ipStr))
		if err != nil {
			return nil, fmt.Errorf("error parsing IP address %s: %w", ipStr, err)
		}
		replicasSlice[i] = fmt.Sprintf("%s:%s", ipAddr, portStr)
	}
	log.Printf("replica IPs: %v", replicasSlice)
	return replicasSlice, nil
}
