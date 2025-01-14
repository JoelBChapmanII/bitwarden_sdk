package sdk

type SecretsInterface interface {
	Create(key, value, note string, organizationID string, projectIDs []string) (*SecretResponse, error)
	List(organizationID string) (*SecretIdentifiersResponse, error)
	Get(secretID string) (*SecretResponse, error)
	Update(secretID string, key, value, note string, organizationID string, projectIDs []string) (*SecretResponse, error)
	Delete(secretIDs []string) (*SecretsDeleteResponse, error)
}

type Secrets struct {
	CommandRunner CommandRunnerInterface
}

func NewSecrets(commandRunner CommandRunnerInterface) *Secrets {
	return &Secrets{CommandRunner: commandRunner}
}

func (s *Secrets) executeCommand(command Command, target interface{}) error {
	responseStr, err := s.CommandRunner.RunCommand(command)
	if err != nil {
		return err
	}
	return checkSuccessAndError(responseStr, target)
}

func (s *Secrets) Create(key, value, note string, organizationID string, projectIDs []string) (*SecretResponse, error) {
	command := Command{
		Secrets: &SecretsCommand{
			Create: &SecretCreateRequest{
				Key:            key,
				Value:          value,
				Note:           note,
				OrganizationID: organizationID,
				ProjectIDS:     projectIDs,
			},
		},
	}

	var response SecretResponse
	if err := s.executeCommand(command, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Secrets) List(organizationID string) (*SecretIdentifiersResponse, error) {
	command := Command{
		Secrets: &SecretsCommand{
			List: &SecretIdentifiersRequest{
				OrganizationID: organizationID,
			},
		},
	}

	var response SecretIdentifiersResponse
	if err := s.executeCommand(command, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Secrets) Get(id string) (*SecretResponse, error) {
	command := Command{
		Secrets: &SecretsCommand{
			Get: &SecretGetRequest{
				ID: id,
			},
		},
	}

	var response SecretResponse
	if err := s.executeCommand(command, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Secrets) Update(id string, key, value, note string, organizationID string, projectIDs []string) (*SecretResponse, error) {
	command := Command{
		Secrets: &SecretsCommand{
			Update: &SecretPutRequest{
				ID:             id,
				Key:            key,
				Value:          value,
				Note:           note,
				OrganizationID: organizationID,
				ProjectIDS:     projectIDs,
			},
		},
	}

	var response SecretResponse
	if err := s.executeCommand(command, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Secrets) Delete(ids []string) (*SecretsDeleteResponse, error) {
	command := Command{
		Secrets: &SecretsCommand{
			Delete: &SecretsDeleteRequest{
				IDS: ids,
			},
		},
	}

	var response SecretsDeleteResponse
	if err := s.executeCommand(command, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
