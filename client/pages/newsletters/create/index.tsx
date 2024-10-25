import { getCookie } from "cookies-next";
import { GetServerSideProps } from "next";
import { Category } from "../../../src/models";
import { getCategories } from "../../../src/adapters/getCategories";
import { withAuth } from "../../../src/components/hocs/withAuth";
import { CreateNewsletterContainer } from "../../../src/components/containers/newsletters/create";

interface NewsletterCreateProps {
  categories: Category[];
}

function NewsletterCreate({ categories }: NewsletterCreateProps) {
  return (
    <main>
      <section className="relative flex items-center justify-center min-h-screen from-purple-900 via-indigo-800 to-indigo-500 bg-gradient-to-br flex-col">
        <CreateNewsletterContainer categories={categories} />
      </section>
    </main>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ req, res }) => {
  const token = getCookie("authToken", { req, res });

  let categories: Category[] = [];

  if (token) {
    const fetchedCategories = await getCategories(token);
    categories = fetchedCategories;
  }

  return { props: { categories } };
};

export default withAuth(NewsletterCreate);
