export interface FastCreationUser {
  email: string;
  password: string;
}

export interface CreateUserDto {
  id: string;
  email: string;
  password: string;
}

export const createUser = async (
  userDto: CreateUserDto
): Promise<FastCreationUser | undefined> => {
  try {
    const response = await fetch(
      `${process.env.USER_SERVICE_API}/v1/users/register`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userDto),
      }
    );

    if (!response.ok) {
      throw new Error("Login failed");
    }

    const data: FastCreationUser = await response.json();

    return data;
  } catch (e) {
    console.error("error creating user ", e);
    return;
  }
};
