[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] | Out-Null
$APP_ID = 'Water time nofifier'
$notification = @"
<toast launch="app-defined-string" duration="long">
  <visual>
    <binding template="ToastGeneric">
      <text>Water timer now!</text>
      <text>Hey, you have worked for an hour, it's better to have some water now.</text>
    </binding>
  </visual>
  <actions>
    <action activationType="foreground" content="Ok, Got it" arguments="accepted"/>
  </actions>
</toast>
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($notification)
$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toast)
