from ibapi.client import EClient
from ibapi.wrapper import EWrapper

from api.IBJts.source.pythonclient.ibapi.account_summary_tags import AccountSummaryTags

class TestWrapper(EWrapper):
    def accountSummary(self, reqId, account, tag, value, currency):
        print("AccountSummary. ReqId:", reqId, "Account:", account, "Tag:", tag, "Value:", value, "Currency:", currency)

    def accountSummaryEnd(self, reqId):
        print("AccountSummaryEnd. ReqId:", reqId)

class TestClient(EClient):
    def __init__(self, wrapper):
        EClient.__init__(self, wrapper)

class TestApp(TestWrapper, TestClient):
    def __init__(self):
        TestWrapper.__init__(self)
        TestClient.__init__(self, wrapper=self)

app = TestApp()
app.connect("192.168.2.100", 7497, clientId=0)
print("here")
app.startApi()  # Start the API connection
app.run()

app.reqAccountSummary(0, "All", AccountSummaryTags.AllTags)
print("here2")

app.disconnect()