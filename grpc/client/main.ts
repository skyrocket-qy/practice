import { createClient, Client } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ListAccountReq } from "./proto/main_pb";
import { Main } from "./proto/main_connect";


const transport = createConnectTransport({
  baseUrl: 'http://localhost:50051', // Use http://localhost:50051
});

const client = createClient(Main, transport);

async function main() {
  try {
      const response = await client.listAccount(new ListAccountReq());
      console.log(response);
  } catch (error) {
      console.error('Error fetching account list:', error);
  }
}

main();