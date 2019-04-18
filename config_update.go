eq := reflect.DeepEqual(existingFluentBitConfigMap, desiredConfigMap)
if !eq {
  reqLogger.Info("FluentBit Config Changed. Updating...")
  existingFluentBitConfigMap = desiredConfigMap
  err = r.client.Update(context.TODO(), existingFluentBitConfigMap)
  reqLogger.Info("FluentBit Config Updated.")
  if err != nil {
    reqLogger.Error(err, "Failed")
  } else {
    err = r.client.Delete(context.TODO(), existingFluentBitDaemonSet)
    return reconcile.Result{Requeue: true}, nil
  }
}
