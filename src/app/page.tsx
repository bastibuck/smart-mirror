import ToDos from "~/apps/ToDos";

export default async function Home() {
  return (
    <main className="bg-dark grid min-h-screen grid-cols-3 grid-rows-3 place-items-stretch p-20">
      <div className="col-start-3 row-start-3 grid place-items-end">
        <ToDos />
      </div>
    </main>
  );
}
