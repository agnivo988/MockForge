"use client";

import { useState } from "react";
import MethodSelect from "./MethodSelect";
import ResponseViewer from "./ResponseViewer";

export default function ApiTester() {
  const [method, setMethod] = useState("GET");
  const [url, setUrl] = useState("/users");
  const [response, setResponse] = useState(null);
  const [status, setStatus] = useState<number | null>(null);

  async function sendRequest() {
    const res = await fetch(`http://localhost:8080${url}`, {
      method,
    });
    const data = await res.json();
    setStatus(res.status);
    setResponse(data);
  }

  return (
    <div className="p-6 grid grid-cols-1 lg:grid-cols-2 gap-6">
      {/* Request Panel */}
      <div className="bg-gray-900 rounded-xl p-4 space-y-4">
        <div className="flex gap-3">
          <MethodSelect value={method} onChange={setMethod} />

          <input
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            className="flex-1 bg-black border border-gray-700 px-3 py-2 rounded text-sm"
            placeholder="/users"
          />

          <button
            onClick={sendRequest}
            className="bg-green-500 text-black px-4 rounded hover:bg-green-400"
          >
            Send
          </button>
        </div>
      </div>

      {/* Response Panel */}
      <div className="bg-gray-900 rounded-xl p-4">
        <ResponseViewer status={status} response={response} />
      </div>
    </div>
  );
}
