import { useEffect, FC, useState } from "react";
import { useRouter } from "next/router";

export const withAuth = <P extends object>(WrappedComponent: FC<P>) => {
  const AuthenticatedComponent: FC<P> = (props) => {
    const router = useRouter();
    const [loading, setLoading] = useState(true);

    useEffect(() => {
      const token = document.cookie
        .split("; ")
        .find((row) => row.startsWith("authToken="));
      if (!token) {
        router.push("/");
      } else {
        setLoading(false);
      }
    }, [router]);

    if (loading) {
      return <p>Loading...</p>;
    }

    return <WrappedComponent {...props} />;
  };

  return AuthenticatedComponent;
};
