/*
 * Copyright (c) Microsoft Corporation. All rights reserved. Licensed under the MIT license.
 * See LICENSE in the project root for license information.
 */
/// <reference path="UIStrings.js" />

'use strict';

(function () {
  var userName = "";
  var email = "dilion.dani@gmail.com";
  if (typeof Office.context === 'undefined') {
    //code for working outside of office client
    //$(document).ready(function () {
    console.log("there is no office context");

    $('#button-text').text("Open Dialog");
    $('#button-desc').text("Open Dialog that shows the workplan");
    $('#run').click(function () {
      run("shotef")
    });
    // document.getElementById('run').addEventListener("click",function() {
    //   run("shotef");
    // });
    $('#ManageTab').click(newPage);
    $('#ShotefTab').click(newPage);
    $('#NewTab').click(newPage);
    $('#HomeTab').click(newPage);
    $('#EmergencyTab').click(newPage);
    //});
  } else {
    // The initialize function must be run each time a new page is loaded
    Office.initialize = function (reason) {
      $(document).ready(function () {
        // Get the language setting for editing document content.
        // To test this, uncomment the following line and then comment out the
        // line that uses Office.context.displayLanguage.
        // var myLanguage = Office.context.contentLanguage;

        // Get the language setting for UI display in the host application.
        var myLanguage = Office.context.displayLanguage;
        userName = Office.context.mailbox.userProfile.displayName;
        email = Office.context.mailbox.userProfile.emailAddress;
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
        $('#action-button').click(openDialogAsIframe);

        $('#ManageTab').click(newPage);
        $('#ShotefTab').click(newPage);
        $('#NewTab').click(newPage);
        $('#HomeTab').click(newPage);
        $('#EmergencyTab').click(newPage);

        $('#run').click(run);
      });
    };


  }



  function run(subject) {
    /**
     * Insert your Outlook code here
     */
    $.ajax({
      // url: "http://localhost:8080/gettasks",
      url: "https://bns-workplan-2018.appspot.com/gettasks",
      headers: {  'X-Appengine-Inbound-Appid': 'bns-workpaln-2018'  },
      method: 'POST',
      dataType: 'json',
      data: {
        email: email,
        subject: subject
      }

    }).done(function (res) {
      console.log(res);
    }).fail(function (error) {
      console.log(error);
    });
    //window.open();

  }

  function openDialogAsIframe(data) {
    //IMPORTANT: IFrame mode only works in Online (Web) clients. Desktop clients (Windows, IOS, Mac) always display as a pop-up inside of Office apps. 

    if (typeof Office.context.ui !== 'undefined') {
      Office.context.ui.displayDialogAsync("https://bnsworkplan.win/dialog", {
        height: 100,
        width: 100,
        displayInIframe: true
      }, null);
    }

  }

  function newPage() {
    if (this.id === 'HomeTab') {
      $('.HomeContent').show().siblings().hide();
    } else if (this.id === 'ShotefTab') {
      $('.ShotefContent').show().siblings().hide();
    } else if (this.id === 'NewTab') {
      $('.NewContent').show().siblings().hide();
    } else if (this.id === 'EmergencyTab') {
      $('.EmergencyContent').show().siblings().hide();
    } else if (this.id === 'ManageTab') {
      $('.ManageContent').show().siblings().hide();
    }
  }


})();