.. Countries FYI documentation master file, created by
   sphinx-quickstart on Mon Mar 13 23:24:15 2017.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

==================
Countries.FYI API!
==================

Deadly simple REST API for getting info about world countries.

GET **v1/countries**
====================

Returns a list of [country_code, country_name] pairs. Designed to use for
autocomplete calls.
Returns a list of country objects if ``full`` flag was specified.

URL
---

https://countries.fyi/v1/countries

Paramenters
-----------

+------+----------+------------------------------------------------------+
| Name | Required | Description                                          |
+======+==========+======================================================+
| full | optional | If specified list of country object will be returned |
+------+----------+------------------------------------------------------+

Request example
---------------

GET https://countries.fyi/v1/countries

Response example
----------------

.. code-block:: javascript

   [
     ["AW","Aruba"],
     ["AF","Afghanistan"],
     ["AO","Angola"],
     // And so on
   ]

GET **v1/countries/:country_id**
================================

Returns an country object by its id.
Two or three letters country code or country name can be used as country id.

URL
---

https://countries.fyi/v1/countries/:country_id

Errors
------

In case there was no country found **Not Found 404** will be returned.

Request example
---------------

GET https://countries.fyi/v1/countries/ua  or

GET https://countries.fyi/v1/countries/ukr or

GET https://countries.fyi/v1/countries/ukraine

Response example
----------------

.. code-block:: json

   {
     "name":{
       "common": "Ukraine",
       "official": "Ukraine",
       "native": {
         "rus": {
           "common": "Украина",
           "official": "Украина"
         },
         "ukr": {
           "common": "Україна",
           "official": "Україна"
         }
       }
     },
     "tld": [".ua", ".укр"],
     "cca2": "UA",
     "ccn3": "804",
     "cca3": "UKR",
     "cioc": "UKR",
     "currency": ["UAH"],
     "callingCode": ["380"],
     "capital": "Kiev",
     "altSpellings": ["UA", "Ukrayina"],
     "region": "Europe",
     "subregion": "Eastern Europe",
     "languages": {
       "rus": "Russian",
       "ukr": "Ukrainian"
     },
     "translations": {
       "deu": {"common": "Ukraine", "official": "Ukraine"},
       "fin": {"common": "Ukraina", "official": "Ukraina"},
       "fra": {"common": "Ukraine", "official": "Ukraine"},
       "hrv": {"common":"Ukrajina", "official": "Ukrajina"},
       "ita": {"common":"Ucraina", "official": "Ucraina"},
       "jpn": {"common":"ウクライナ", "official": "ウクライナ"},
       "nld": {"common":"Oekraïne", "official": "Oekraïne"},
       "por": {"common":"Ucrânia", "official": "Ucrânia"},
       "rus": {"common":"Украина", "official": "Украина"},
       "slk": {"common":"Ukrajina", "official": "Ukrajina"},
       "spa": {"common":"Ucrania", "official": "Ucrania"},
       "zho": {"common":"乌克兰", "official": "乌克兰"}
     },
     "latlng": [49, 32],
     "denonym": "",
     "landlocked": false,
     "borders": ["BLR", "HUN", "MDA", "POL", "ROU", "RUS", "SVK"],
     "area": 603500
   }

Sources
-------

Based on countries data from https://github.com/mledoze/countries
