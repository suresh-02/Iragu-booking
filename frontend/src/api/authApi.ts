import createApiSession from "./api";

interface CreateUserPayload {
  username: string;
  email: string;
  password: string;
}

interface loginUserPayload {
  email: string;
  password: string;
}

export const createUser = async (payload: CreateUserPayload) => {
  try {
    const response = await createApiSession.post("/register", payload);
    return response.data;
  } catch (error) {
    const response = "Failed to create user";
    return response;
  }
};

export const loginUser = async (payload: loginUserPayload) => {
  try {
    const response = await createApiSession.post("/login", payload);
    return response.data;
  } catch (error) {
    const response = "Failed to login user";
    return response;
  }
};
