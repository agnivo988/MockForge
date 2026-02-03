"use client";

import { useState } from "react";

export default function HeadersEditor({
  onChange,
}: {
  onChange: (h: Record<string, string>) => void;
}) {
  const [key, setKey] = useState("");
  const [value, setValue] = useState("");

  function addHeader() {
    onChange({ [key]: value });
    setKey("");
    setValue("");
  }

  return (
    <div>
      <p className="text-sm font-semibold">Headers</p>
      <div className="flex gap-2 mt-1">
        <input
          placeholder="Key"
          value={key}
          onChange={(e) => setKey(e.target.value)}
          className="border px-2 py-1"
        />
        <input
          placeholder="Value"
          value={value}
          onChange={(e) => setValue(e.target.value)}
          className="border px-2 py-1"
        />
        <button
          onClick={addHeader}
          className="border px-3"
        >
          Add
        </button>
      </div>
    </div>
  );
}
