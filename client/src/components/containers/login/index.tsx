import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { createUser, CreateUserDto } from "../../../adapters/createUser";
import { login } from "../../../adapters/authUser";
import { DEFAULT_ADMIN } from "../../../constants/admin";
import { CREATE_USER_KEY } from "../../../constants/localStorage";
import {
  checkLocalStorageValueExists,
  getLocalStorageValue,
  setCookie,
  setLocalStorageValue,
} from "../../../utils";
import { LoginComponent } from "../../lib/Login.component";
import { NotificationComponent } from "../../lib/Notification/Notification.component";
import { ToastComponent } from "../../lib/Notification/Toast.component";

export default function LoginContainer() {
  const router = useRouter();
  const [isUserCreationLoading, setIsUserCreationLoading] = useState(true);
  const [isUserCreated, setIsUserCreated] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  useEffect(() => {
    const isFastCreationUserExist =
      checkLocalStorageValueExists(CREATE_USER_KEY);

    setFormValues(isFastCreationUserExist);
    setIsUserCreated(isFastCreationUserExist);
    setIsUserCreationLoading(false);
  }, []);

  const setFormValues = (exist: boolean) => {
    if (exist) {
      const user = getLocalStorageValue<CreateUserDto>(CREATE_USER_KEY);
      if (user?.email) {
        setEmail(user.email);
      }

      if (user?.password) {
        setPassword(user.password);
      }
    }
  };

  const handleCreateUser = async () => {
    const userDto: CreateUserDto = {
      id: DEFAULT_ADMIN.id,
      email: DEFAULT_ADMIN.email,
      password: DEFAULT_ADMIN.password,
    };

    setIsUserCreationLoading(true);
    const user = await createUser(userDto);
    if (!user) {
      // Error
      setIsUserCreationLoading(false);
      return;
    }
    setLocalStorageValue(CREATE_USER_KEY, userDto);

    setFormValues(true);
    setIsUserCreated(true);
  };

  const handleAuthUser = async () => {
    const response = await login({ email, password });
    if (!response?.token) {
      // Error
      return;
    }

    setCookie("authToken", response.token);
    router.replace("/newsletters");
  };

  return (
    <main>
      <section className="relative">
        <LoginComponent
          email={email}
          password={password}
          onClick={handleAuthUser}
        />
      </section>

      <aside className="absolute" style={{ top: "5%", right: "2%" }}>
        {isUserCreated ? (
          <NotificationComponent />
        ) : (
          <ToastComponent
            onClick={handleCreateUser}
            isLoading={isUserCreationLoading}
          />
        )}
      </aside>
    </main>
  );
}
