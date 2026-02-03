"use client";

import { useState } from "react";
import MethodSelect from "./MethodSelect";
import HeadersEditor from "./HeadersEditor";
import ResponseViewer from "./ResponseViewer";

export default function RequestBuilder() {
  const [method, setMethod] = useState("GET");
  const [url, setUrl] = useState("http://localhost:8080/");
  const [headers, setHeaders] = useState<Record<string, string>>({});
  const [body, setBody] = useState("");
  const [response, setResponse] = useState<any>(null);
  const [status, setStatus] = useState<number | null>(null);
  const [activeTab, setActiveTab] = useState<"headers" | "body">("headers");
  const [loading, setLoading] = useState(false);

  async function sendRequest() {
    try {
      setLoading(true);
      const res = await fetch(url, {
        method,
        headers: {
          "Content-Type": "application/json",
          ...headers,
        },
        body: method !== "GET" ? body : undefined,
      });

      const data = await res.json();
      setStatus(res.status);
      setResponse(data);
    } catch (err) {
      setStatus(0);
      setResponse({ error: "Failed to fetch" });
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="bg-gray-900 text-white rounded-xl p-5 space-y-4 shadow-lg">
      {/* Top Bar */}
      <div className="flex gap-2 items-center">
        <MethodSelect value={method} onChange={setMethod} />

        <input
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          className="flex-1 bg-black border border-gray-700 px-3 py-2 rounded text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="http://localhost:8080/users"
        />

        <button
          onClick={sendRequest}
          disabled={loading}
          className="bg-blue-500 hover:bg-blue-400 text-black px-5 py-2 rounded font-medium disabled:opacity-50"
        >
          {loading ? "Sending..." : "Send"}
        </button>
      </div>

      {/* Tabs */}
      <div className="flex gap-4 border-b border-gray-700 text-sm">
        <button
          onClick={() => setActiveTab("headers")}
          className={`pb-2 ${
            activeTab === "headers"
              ? "border-b-2 border-blue-500 text-blue-400 cursor-pointer"
              : "text-gray-400 cursor-pointer"
          }`}
        >
          Headers
        </button>

        {method !== "GET" && (
          <button
            onClick={() => setActiveTab("body")}
            className={`pb-2 ${
              activeTab === "body"
                ? "border-b-2 border-blue-500 text-blue-400 cursor-pointer"
                : "text-gray-400 cursor-pointer"
            }`}
          >
            Body
          </button>
        )}
      </div>

      {/* Tab Content */}
      <div>
        {activeTab === "headers" && (
          <HeadersEditor onChange={setHeaders} />
        )}

        {activeTab === "body" && method !== "GET" && (
          <textarea
            className="w-full h-40 bg-black border border-gray-700 p-3 rounded font-mono text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder='{"key": "value"}'
            value={body}
            onChange={(e) => setBody(e.target.value)}
          />
        )}
      </div>

      {/* Response */}
      <div className="pt-2">
        <ResponseViewer status={status} response={response} />
      </div>
    </div>
  );
}
