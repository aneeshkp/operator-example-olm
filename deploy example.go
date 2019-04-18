func newDaemonset(cr *v1beta1.GenericDaemon) *appsv1.DaemonSet {
  return &appsv1.DaemonSet{
    TypeMeta: metav1.TypeMeta{
      Kind:       "DaemonSet",
      APIVersion: "apps/v1",
    },
    ObjectMeta: metav1.ObjectMeta{
      Name:      cr.Name + "-daemonset",
      Namespace: cr.Namespace,
      OwnerReferences: []metav1.OwnerReference{
        *metav1.NewControllerRef(cr, schema.GroupVersionKind{
          Group:   v1beta1.SchemeGroupVersion.Group,
          Version: v1beta1.SchemeGroupVersion.Version,
          Kind:    "GenericDaemon",
        }),
      },
    },
    Spec: appsv1.DaemonSetSpec{
      Selector: &metav1.LabelSelector{
        MatchLabels: map[string]string{"daemonset": cr.Name + "-daemonset"},
      },
      Template: corev1.PodTemplateSpec{
        ObjectMeta: metav1.ObjectMeta{
          Labels: map[string]string{"daemonset": cr.Name + "-daemonset"},
        },
        Spec: corev1.PodSpec{
          NodeSelector: map[string]string{"daemon": cr.Spec.Label},
          Containers: []corev1.Container{
            {
              Name:  "genericdaemon",
              Image: cr.Spec.Image,
            },
          },
        },
      },
    },
  }
