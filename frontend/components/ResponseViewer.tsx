export default function ResponseViewer({
  status,
  response,
}: {
  status: number | null;
  response: any;
}) {
  if (!response) return null;

  return (
    <div className="bg-black p-3 rounded text-sm text-white font-mono flex flex-col mt-4">
      <p className="mb-2">
        Status: <strong>{status}</strong>
      </p>

      <pre className="
        overflow-x-auto
        overflow-y-auto
        max-h-96
        whitespace-pre-wrap
        break-words
      ">
        {JSON.stringify(response, null, 2)}
      </pre>
    </div>
  );
}

