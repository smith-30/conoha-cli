package model

import "time"

type (
	AuthRes struct {
		Access struct {
			Metadata struct {
				IsAdmin int      `json:"is_admin"`
				Roles   []string `json:"roles"`
			} `json:"metadata"`
			ServiceCatalog []struct {
				Endpoints []struct {
					PublicURL string `json:"publicURL"`
					Region    string `json:"region"`
				} `json:"endpoints"`
				EndpointsLinks []interface{} `json:"endpoints_links"`
				Name           string        `json:"name"`
				Type           string        `json:"type"`
			} `json:"serviceCatalog"`
			Token struct {
				AuditIds []string  `json:"audit_ids"`
				Expires  time.Time `json:"expires"`
				ID       string    `json:"id"`
				IssuedAt string    `json:"issued_at"`
				Tenant   struct {
					Description string `json:"description"`
					DomainID    string `json:"domain_id"`
					Enabled     bool   `json:"enabled"`
					ID          string `json:"id"`
					Name        string `json:"name"`
				} `json:"tenant"`
			} `json:"token"`
			User struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Roles []struct {
					Name string `json:"name"`
				} `json:"roles"`
				RolesLinks []interface{} `json:"roles_links"`
				Username   string        `json:"username"`
			} `json:"user"`
		} `json:"access"`
	}
)
