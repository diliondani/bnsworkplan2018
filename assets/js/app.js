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

            // Set localized text for UI elements.
            $("h1").text(userName + " " +  UIText.Greeting );
            $("#about").text(UIText.Introduction);
      $('#run').click(run);
    });
  };

  function run() {
    
    
    /**
     * Insert your Outlook code here
     */
    
  }

})();