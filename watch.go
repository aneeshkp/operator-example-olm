func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("logmanagement-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

  // Watch for the primary LogManagement Resource
	err = c.Watch(&source.Kind{Type: &loggingv1alpha1.LogManagement{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

  // Watch for deployment
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &loggingv1alpha1.LogManagement{},
	})

  // Watch for configmap
	err = c.Watch(&source.Kind{Type: &corev1.ConfigMap{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &loggingv1alpha1.LogManagement{},
	})
  
  // Watch for daemonset
	err = c.Watch(&source.Kind{Type: &extensionv1.DaemonSet{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &loggingv1alpha1.LogManagement{},
	})

  // Watch for service
	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &loggingv1alpha1.LogManagement{},
	})

	return nil
}
