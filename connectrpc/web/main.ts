import { createClient, Client } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { Main } from "./proto/main_connect"; // Adjust based on your generated output
import { ListAccountReq } from "./proto/main_pb"; // Adjust as needed

// Set up the transport for connect-web
const transport = createConnectTransport({
  baseUrl: "http://localhost:50051", // Set to the URL of your proxy server (e.g., Envoy or grpc-web)
});

// Create the client for the Main service
const mainClient = createClient(Main, transport);

// Function to call listAccount
export async function fetchAccounts() {
  try {
    const request = new ListAccountReq();
    const response = await mainClient.listAccount(request);
    console.log("Accounts fetched:", response);
    return response;
  } catch (error) {
    console.error("Error fetching accounts:", error);
    throw error;
  }
}

// Call fetchAccounts to test
fetchAccounts().catch(console.error);
