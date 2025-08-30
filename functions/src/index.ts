import * as functions from "firebase-functions";
import * as admin from "firebase-admin";

admin.initializeApp();

export const api = functions.https.onRequest((request, response) => {
  response.set("Access-Control-Allow-Origin", "*");
  response.set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS");
  response.set("Access-Control-Allow-Headers", "Content-Type, Authorization");

  if (request.method === "OPTIONS") {
    response.status(204).send("");
    return;
  }

  // Proxy to your Go backend or implement API logic here
  response.json({
    message: "API endpoint - implement your backend logic here",
    path: request.path,
    method: request.method
  });
});