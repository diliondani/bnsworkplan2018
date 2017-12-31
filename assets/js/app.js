/*
 * Copyright (c) Microsoft Corporation. All rights reserved. Licensed under the MIT license.
 * See LICENSE in the project root for license information.
 */
/// <reference path="UIStrings.js" />

'use strict';

(function () {

  // The initialize function must be run each time a new page is loaded
  Office.initialize = function (reason) {
    $(document).ready(function () {
      // Get the language setting for editing document content.
      // To test this, uncomment the following line and then comment out the
      // line that uses Office.context.displayLanguage.
      // var myLanguage = Office.context.contentLanguage;

      // Get the language setting for UI display in the host application.
      var myLanguage = Office.context.displayLanguage;
      var userName = Office.context.mailbox.userProfile.displayName;
      //var email = Office.context.mailbox.userProfile.emailAddress;
      var UIText;

      // Get the resource strings that match the language.
      // Use the UIStrings object from the UIStrings.js file
      // to get the JSON object with the correct localized strings.
      UIText = UIStrings.getLocaleStrings(myLanguage);

      // // Initialize the FabricUI notification mechanism and hide it
      // var element = document.querySelector('.ms-MessageBanner');
      // messageBanner = new app.notification.MessageBanner(element);
      // messageBanner.hideBanner();

      // Set localized text for UI elements.
      $("h1").text(userName + " " + UIText.Greeting);
      $("#about").text(UIText.Introduction);

      // Initialize fabric components
      //$('.ms-ContextualMenu').ContextualMenu();

      $('#button-text').text("Open Dialog");
      $('#button-desc').text("Open Dialog that shows the workplan");
      $('#action-button').click(openDialogAsIframe)

      $('#run').click(run);
    });
  };

  // var dialog;
  // // Display notifications in message banner at the top of the task pane.
  // function showNotification(content) {
  //   $("#notificationBody").text(content);
  //   messageBanner.showBanner();
  //   messageBanner.toggleExpansion();
  // }

  function run() {


    /**
     * Insert your Outlook code here
     */

  }

  // function dialogCallback(asyncResult) {
  //   if (asyncResult.status == "failed") {

  //     // In addition to general system errors, there are 3 specific errors for 
  //     // displayDialogAsync that you can handle individually.
  //     switch (asyncResult.error.code) {
  //       case 12004:
  //         showNotification("Domain is not trusted");
  //         break;
  //       case 12005:
  //         showNotification("HTTPS is required");
  //         break;
  //       case 12007:
  //         showNotification("A dialog is already opened.");
  //         break;
  //       default:
  //         showNotification(asyncResult.error.message);
  //         break;
  //     }
  //   } else {
  //     dialog = asyncResult.value;
  //     /*Messages are sent by developers programatically from the dialog using office.context.ui.messageParent(...)*/
  //     dialog.addEventHandler(Office.EventType.DialogMessageReceived, messageHandler);

  //     /*Events are sent by the platform in response to user actions or errors. For example, the dialog is closed via the 'x' button*/
  //     dialog.addEventHandler(Office.EventType.DialogEventReceived, eventHandler);
  //   }
  // }

  function openDialogAsIframe() {
    //IMPORTANT: IFrame mode only works in Online (Web) clients. Desktop clients (Windows, IOS, Mac) always display as a pop-up inside of Office apps. 

    Office.context.ui.displayDialogAsync("https://bnsworkplan.win/assets/html/Dialog.html", {
      height: 50,
      width: 50,
      displayInIframe: true
    }, null);
  }


})();