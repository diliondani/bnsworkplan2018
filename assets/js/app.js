/*
 * Copyright (c) Microsoft Corporation. All rights reserved. Licensed under the MIT license.
 * See LICENSE in the project root for license information.
 */
/// <reference path="UIStrings.js" />
/// <reference path="fabric.js" />

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
      $("h1").text(userName + " " + UIText.Greeting);
      $("#about").text(UIText.Introduction);

      // Initialize fabric components
      //$('.ms-ContextualMenu').ContextualMenu();


      $('#myButton').click(function () {
        document.getElementById("myDropdown").classList.toggle("is-open");
      });

      var DialogComponents = [];
      var DialogElements = $('.ms-Dialog');
      for (var i = 0; i < DialogElements.length; i++) {
        (function() {
          DialogComponents[i] = new fabric['Dialog'](DialogElements[i]);
        }());
      }
      $('#openDialog').click(function (event){
        $('.ms-Dialog').open()
      })
      $('#run').click(run);
    });
  };

  function run() {


    /**
     * Insert your Outlook code here
     */

  }

  function ContextualMenu() {
    console.log("ContextualMenu was called")

    /** Go through each contextual menu we've been given. */
    return this.each(function () {

      var $contextualMenu = $(this);

      // Set selected states.
      $contextualMenu.on('click', '.ms-ContextualMenu-link:not(.is-disabled)', function (event) {
        event.preventDefault();

        // Check if multiselect - set selected states
        if ($contextualMenu.hasClass('ms-ContextualMenu--multiselect')) {

          // If already selected, remove selection; if not, add selection
          if ($(this).hasClass('is-selected')) {
            $(this).removeClass('is-selected');
          } else {
            $(this).addClass('is-selected');
          }

        }
        // All other contextual menu variants
        else {

          // Deselect all of the items and close any menus.
          $('.ms-ContextualMenu-link')
            .removeClass('is-selected')
            .siblings('.ms-ContextualMenu')
            .removeClass('is-open');

          // Select this item.
          $(this).addClass('is-selected');

          // If this item has a menu, open it.
          if ($(this).hasClass('ms-ContextualMenu-link--hasMenu')) {
            $(this).siblings('.ms-ContextualMenu:first').addClass('is-open');

            // Open the menu without bubbling up the click event,
            // which can cause the menu to close.
            event.stopPropagation();
          }

        }
      });
    });
  };

})();