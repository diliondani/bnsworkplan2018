/* Store the locale-specific strings */

var UIStrings = (function ()
{
    "use strict";

    var UIStrings = {};

    // JSON object for English strings
    UIStrings.EN =
    {        
        "Greeting": "Welcome",
        "Introduction": "This is my localized add-in."        
    };

    // JSON object for Hebrew strings
    UIStrings.HE =
    {        
        "Greeting": "שלום",
        "Introduction": "תכנית העבודה לשנת 2018"
    };

    UIStrings.getLocaleStrings = function (locale)
    {
        var text;

        // Get the resource strings that match the language.
        switch (locale)
        {
            case 'en-US':
                text = UIStrings.EN;
                break;
            case 'he-IL':
                text = UIStrings.HE;
                break;
            default:
                text = UIStrings.EN;
                break;
        }

        return text;
    };

    return UIStrings;
})();