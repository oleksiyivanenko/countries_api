.. Countries FYI documentation master file, created by
   sphinx-quickstart on Mon Mar 13 23:24:15 2017.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

==================
Countries FYI API!
==================

Deadly simple API for getting info about countries.

Methods
-------

GET ``v1/countries``

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

Real method url https://countries.fyi/v1/countries

**GET** ``v1/countries/<country_code>``

For example https://countries.fyi/v1/countries/ua
