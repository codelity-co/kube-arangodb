package backup

import (
	"fmt"
	database "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1alpha"
	"github.com/arangodb/kube-arangodb/pkg/backup/operator"
	"github.com/arangodb/kube-arangodb/pkg/backup/state"
	fakeClientSet "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/fake"
	"github.com/stretchr/testify/require"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"testing"
)

func newFakeHandler() *handler {
	f := fakeClientSet.NewSimpleClientset()

	return &handler{
		client:f,
		arangoClientTimeout:defaultArangoClientTimeout,
	}
}

func newErrorsFakeHandler(errors mockErrorsArangoClientBackup) (*handler,  *mockArangoClientBackup) {
	handler := newFakeHandler()

	mock := newMockArangoClientBackup(errors)
	handler.arangoClientFactory = newMockArangoClientBackupFactory(mock)

	return handler, mock
}

func newObjectSet(state state.State) (*database.ArangoBackup, *database.ArangoDeployment) {
	name := string(uuid.NewUUID())
	namespace := string(uuid.NewUUID())

	obj := newArangoBackup(name, namespace, name, state)
	deployment := newArangoDeployment(namespace, name)

	return obj, deployment
}

func compareTemporaryState(t *testing.T, err error, errorMsg string, handler *handler, obj *database.ArangoBackup) {
	require.Error(t, err)
	require.True(t, IsTemporaryError(err))
	require.EqualError(t, err, errorMsg)

	newObj := refreshArangoBackup(t, handler, obj)
	require.Equal(t, obj.Status, newObj.Status)
}

func newItem(operation operator.Operation, namespace, name string) operator.Item {
	return operator.Item{
		Group: database.SchemeGroupVersion.Group,
		Version: database.SchemeGroupVersion.Version,
		Kind:database.ArangoBackupResourceKind,

		Operation:operation,

		Namespace:namespace,
		Name:name,
	}
}

func newItemFromBackup(operation operator.Operation, backup *database.ArangoBackup) operator.Item {
	return newItem(operation, backup.Namespace, backup.Name)
}

func newArangoBackup(objectRef, namespace, name string, state state.State) *database.ArangoBackup {
	return &database.ArangoBackup{
		TypeMeta: meta.TypeMeta{
			APIVersion: database.SchemeGroupVersion.String(),
			Kind: database.ArangoBackupResourceKind,
		},
		ObjectMeta: meta.ObjectMeta{
			Name: name,
			Namespace:namespace,
			SelfLink: fmt.Sprintf("/api/%s/%s/%s/pods/%s",
				database.SchemeGroupVersion.String(),
				database.ArangoBackupResourcePlural,
				namespace,
				name),
			UID: uuid.NewUUID(),
		},
		Spec:database.ArangoBackupSpec{
			Deployment: database.ArangoBackupSpecDeployment{
				Name: objectRef,
			},
		},
		Status:database.ArangoBackupStatus{
			State:database.ArangoBackupState{
				State: state,
			},
		},
	}
}

func createArangoBackup(t *testing.T, h *handler, backups ... *database.ArangoBackup) {
	for _, backup := range backups {
		_, err := h.client.DatabaseV1alpha().ArangoBackups(backup.Namespace).Create(backup)
		require.NoError(t, err)
	}
}

func refreshArangoBackup(t *testing.T, h *handler, backup *database.ArangoBackup) *database.ArangoBackup{
	obj, err := h.client.DatabaseV1alpha().ArangoBackups(backup.Namespace).Get(backup.Name, meta.GetOptions{})
	require.NoError(t, err)
	return obj
}

func newArangoDeployment(namespace, name string) *database.ArangoDeployment {
	return &database.ArangoDeployment{
		TypeMeta: meta.TypeMeta{
			APIVersion: database.SchemeGroupVersion.String(),
			Kind: database.ArangoDeploymentResourceKind,
		},
		ObjectMeta: meta.ObjectMeta{
			Name: name,
			Namespace:namespace,
			SelfLink: fmt.Sprintf("/api/%s/%s/%s/pods/%s",
				database.SchemeGroupVersion.String(),
				database.ArangoDeploymentResourcePlural,
				namespace,
				name),
			UID: uuid.NewUUID(),
		},
	}
}

func createArangoDeployment(t *testing.T, h *handler, deployments ... *database.ArangoDeployment) {
	for _, deployment := range deployments {
		_, err := h.client.DatabaseV1alpha().ArangoDeployments(deployment.Namespace).Create(deployment)
		require.NoError(t, err)
	}
}

func wrapperUndefinedDeployment(t *testing.T, state state.State) {
	t.Run("Empty Name", func(t *testing.T) {
		// Arrange
		handler := newFakeHandler()

		obj, _ := newObjectSet(state)
		obj.Spec.Deployment.Name = ""

		// Act
		createArangoBackup(t, handler, obj)
		require.NoError(t, handler.Handle(newItemFromBackup(operator.OperationUpdate, obj)))

		// Assert
		newObj := refreshArangoBackup(t, handler, obj)
		require.Equal(t, newObj.Status.State.State, database.ArangoBackupStateFailed)

		require.Equal(t, newObj.Status.State.Message, fmt.Sprintf("deployment ref is not specified for backup %s/%s", obj.Namespace, obj.Name))
	})

	t.Run("Missing Deployment", func(t *testing.T) {
		// Arrange
		handler := newFakeHandler()

		obj, _ := newObjectSet(state)

		// Act
		createArangoBackup(t, handler, obj)
		require.NoError(t, handler.Handle(newItemFromBackup(operator.OperationUpdate, obj)))

		// Assert
		newObj := refreshArangoBackup(t, handler, obj)
		require.Equal(t, newObj.Status.State.State, database.ArangoBackupStateFailed)

		require.Equal(t, newObj.Status.State.Message, fmt.Sprintf("%s \"%s\" not found", database.ArangoDeploymentCRDName, obj.Name))
	})
}

func wrapperConnectionIssues(t *testing.T, state state.State) {
	t.Run("Unable to create deployment client", func(t *testing.T) {
		// Arrange
		handler := newFakeHandler()

		f := newMockArangoClientBackupErrorFactory(fmt.Errorf("error"))
		handler.arangoClientFactory = f

		obj, deployment := newObjectSet(state)

		// Act
		createArangoBackup(t, handler, obj)
		createArangoDeployment(t, handler, deployment)
		err := handler.Handle(newItemFromBackup(operator.OperationUpdate, obj))

		// Assert
		require.Error(t, err)
		require.True(t, IsTemporaryError(err))

		newObj := refreshArangoBackup(t, handler, obj)
		require.Equal(t, newObj.Status.State.State, state)

	})
}

func wrapperProgressMissing(t *testing.T, state state.State) {
	t.Run("Backup Progress Missing", func(t *testing.T) {
		// Arrange
		handler, _ := newErrorsFakeHandler(mockErrorsArangoClientBackup{})

		obj, deployment := newObjectSet(state)

		// Act
		createArangoBackup(t, handler, obj)
		createArangoDeployment(t, handler, deployment)
		require.NoError(t, handler.Handle(newItemFromBackup(operator.OperationUpdate, obj)))

		// Assert
		newObj := refreshArangoBackup(t, handler, obj)
		require.Equal(t, newObj.Status.State.State, database.ArangoBackupStateFailed)

		require.Equal(t, newObj.Status.State.Message, fmt.Sprintf("backup details are missing"))

	})
}