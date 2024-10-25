import { getCookie } from "cookies-next";
import { useRouter } from "next/router";
import { GetServerSideProps } from "next";

import { Newsletter } from "../../src/models";
import { getNewsletters } from "../../src/adapters/getNewsletter";
import { withAuth } from "../../src/components/hocs/withAuth";
import { NewsletterContainer } from "../../src/components/containers/newsletters";

interface NewslettersPageProps {
  newsletters: Newsletter[];
}

function Newsletters({ newsletters }: NewslettersPageProps) {
  const router = useRouter();

  const handleRedirectToCreateNewsletter = () => {
    router.push("/newsletters/create");
  };

  return (
    <main>
      <section className="relative flex items-center justify-center min-h-screen from-purple-900 via-indigo-800 to-indigo-500 bg-gradient-to-br flex-col">
        <NewsletterContainer newsletters={newsletters} />
      </section>
      <aside style={{ position: "absolute", top: "93%", right: "1%" }}>
        <button
          onClick={handleRedirectToCreateNewsletter}
          type="button"
          className="text-white bg-gradient-to-r from-purple-500 via-purple-600 to-purple-800 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-purple-300 dark:focus:ring-purple-800 shadow-lg shadow-purple-500/50 dark:shadow-lg dark:shadow-purple-800/80 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
        >
          Create Newsletter
        </button>
      </aside>
    </main>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ req, res }) => {
  const token = getCookie("authToken", { req, res });

  let newsletters: Newsletter[] = [];

  if (token) {
    const fetchedNewsletters = await getNewsletters(token);
    newsletters = fetchedNewsletters;
  }

  return { props: { newsletters } };
};

export default withAuth(Newsletters);
