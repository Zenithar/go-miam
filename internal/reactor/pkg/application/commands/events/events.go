package events

// ApplicationCreated is raised when an application domain has been created.
func ApplicationCreated(id, label string) interface{} {
	return nil
}

// ApplicationActivated is raised when an application domain has been activated.
func ApplicationActivated(id string) interface{} {
	return nil
}

// ApplicationDisabled is raised when an application domain has been disabled.
func ApplicationDisabled(id string) interface{} {
	return nil
}

// ApplicationLabelChanged is raised when an application label attribute has been changed.
func ApplicationLabelChanged(id, old, new string) interface{} {
	return nil
}
