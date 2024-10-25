import Head from "next/head";
import React from "react";

interface LayoutProps {
  children: React.ReactNode; // Define the type for children
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <>
      <Head>
        <title>Stori Challenge</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>{children}</main>
    </>
  );
};

export { Layout };
