![CRUDA - e2e test](logo-h80.png "CRUDA e2e")

E2E Tests for **[CRUDA](http://github.com/Oppodelldog/cruda)**

These tests written in go utilize **[chromedp](http://github.com/chromedp/chromedp)** which integrates  
directly with chrome implementing the dev tools protocol.

Since **CRUDA** itself is a lock-down-weekend-fun project and so its  
scope is small it was fun to set up these tests to cover the use cases.  

#### Project Structure

|Folder|Description|
|---|---|
|**cmd**|the main entry point defines and runs the test suites|
|**cases**|contains the test cases|
|**cruda**|project specific high level testing API|
