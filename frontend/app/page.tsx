import RequestBuilder from "@/components/RequestBuilder";

export default function Home() {
  return (
    <main className="min-h-screen p-10">
      <h1 className="text-3xl font-bold mb-4">MockForge</h1>
      <RequestBuilder />
    </main>
  );
}
