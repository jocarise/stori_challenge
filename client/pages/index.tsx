import dynamic from "next/dynamic";

const NoSSR = dynamic(() => import("../src/components/containers/login"), {
  ssr: false,
});

function Home() {
  return <NoSSR />;
}

export default Home;
