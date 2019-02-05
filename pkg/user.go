package pkg

import "errors"

type AuthProviderConfig struct {
	AccessToken string `yaml:"access-token,omitempty"`
	CmdArgs     string `yaml:"cmd-args,omitempty"`
	CmdPath     string `yaml:"cmd-path,omitempty"`
	Expiry      string `yaml:"expiry,omitempty"`
	ExpiryKey   string `yaml:"expiry-key,omitempty"`
	TokenKey    string `yaml:"token-key,omitempty"`
}

type AuthProvider struct {
	Config AuthProviderConfig `yaml:"config"`
	Name   string             `yaml:"name"`
}

type UserConf struct {
	AuthProvider      AuthProvider `yaml:"auth-provider,omitempty"`
	Password          string       `yaml:"password,omitempty"`
	Username          string       `yaml:"username,omitempty"`
	Token             string       `yaml:"token,omitempty"`
	ClientCertificate string       `yaml:"client-certificate,omitempty"`
	ClientKey         string       `yaml:"client-key,omitempty"`
}

type User struct {
	Name     string   `yaml:"name"`
	UserConf UserConf `yaml:"user"`
}

type Users []User

func (u Users) Len() int           { return len(u) }
func (u Users) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u Users) Less(i, j int) bool { return u[i].Name < u[j].Name }

func (u Users) Get(name string) (User, error) {
	for _, user := range u {
		if user.Name == name {
			return user, nil
		}
	}
	return User{}, errors.New("no user for name: " + name)
}

type UserFilter func(ctx User) bool

func (u Users) Filter(filter UserFilter) Users {
	contexts := make(Users, 0, len(u))
	for _, ctx := range u {
		if filter(ctx) {
			contexts = append(contexts, ctx)
		}
	}
	return contexts
}
