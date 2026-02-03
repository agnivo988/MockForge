"use client";

export default function UploadBox() {
  return (
    <div className="mt-6 p-6 border rounded-lg">
      <h2 className="text-lg font-semibold">Upload OpenAPI Spec</h2>
      <input
        type="file"
        className="mt-4"
        accept=".yaml,.yml,.json"
      />
      <p className="text-sm text-gray-500 mt-2">
        (Backend hot-reload coming next)
      </p>
    </div>
  );
}
