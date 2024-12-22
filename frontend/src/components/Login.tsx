import { useState } from "react";
import { Eye, EyeOff } from "lucide-react";
import loginImage from "../assets/loginImage.svg";
import { createUser, loginUser } from "../api/authApi";

const Login = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
  });
  const [message, setMessage] = useState("");
  const [signUp, setSignUp] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (signUp) {
        const response = await createUser(formData);
        setMessage(() => response.message);
        return;
      }
      const response = await loginUser(formData);
      setMessage(() => response.message);
      return;
    } catch (err) {
      setMessage(() => "Error Occured");
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const isSignIn = (event: React.MouseEvent<HTMLAnchorElement, MouseEvent>) => {
    event.preventDefault();
    setSignUp((prevState) => !prevState);
    return signUp;
  };

  return (
    <div className="min-h-screen flex flex-col lg:flex-row">
      {/* Left Column - Image */}
      <div className="w-1/2 bg-green-50 hidden md:block">
        <div className="h-full flex items-center justify-center p-8">
          <img
            src={loginImage}
            alt="Badminton Court"
            className="rounded-xl shadow-lg object-cover max-h-[600px]"
          />
        </div>
      </div>

      {/* Right Column - Login Form */}
      <div className="lg:w-1/2 flex items-center justify-center p-8">
        <div className="max-w-md w-full space-y-8">
          <div className="text-center">
            {signUp ? (
              <div>
                <h2 className="text-3xl font-bold text-gray-900">Welcome..!</h2>
                <p className="mt-2 text-gray-600">Please Create your account</p>
              </div>
            ) : (
              <div>
                <h2 className="text-3xl font-bold text-gray-900">
                  Welcome Back..!
                </h2>
                <p className="mt-2 text-gray-600">
                  Please sign in to your account
                </p>
              </div>
            )}
          </div>

          <form onSubmit={handleSubmit} className="mt-8 space-y-6">
            <div className="space-y-4">
              {signUp ? (
                <div>
                  <label
                    htmlFor="email"
                    className="block text-start font-bold text-sm  text-gray-700"
                  >
                    Username
                  </label>
                  <input
                    id="username"
                    name="username"
                    type="text"
                    required
                    value={formData.username}
                    onChange={handleChange}
                    className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    placeholder="Enter your email"
                  />
                </div>
              ) : (
                ""
              )}

              <div>
                <label
                  htmlFor="email"
                  className="block text-start font-bold text-sm  text-gray-700"
                >
                  Email address
                </label>
                <input
                  id="email"
                  name="email"
                  type="email"
                  autoComplete="email"
                  required
                  value={formData.email}
                  onChange={handleChange}
                  className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                  placeholder="Enter your email"
                />
              </div>

              <div className="relative">
                <label
                  htmlFor="password"
                  className="block text-start font-bold text-sm text-gray-700"
                >
                  Password
                </label>
                <div className="mt-1 relative">
                  <input
                    id="password"
                    name="password"
                    type={showPassword ? "text" : "password"}
                    autoComplete="current-password"
                    required
                    value={formData.password}
                    onChange={handleChange}
                    className="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    placeholder="Enter your password"
                  />
                  <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute inset-y-0 right-0 pr-3 flex items-center"
                  >
                    {showPassword ? (
                      <EyeOff className="h-5 w-5 text-gray-400" />
                    ) : (
                      <Eye className="h-5 w-5 text-gray-400" />
                    )}
                  </button>
                </div>
              </div>
            </div>

            <div className="flex items-center justify-between">
              <div className="text-sm">
                <a
                  href="#"
                  className="font-medium text-green-600 hover:text-green-500"
                >
                  Forgot your password?
                </a>
              </div>
            </div>

            <div>
              {signUp ? (
                <button
                  type="submit"
                  className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-400 hover:bg-green-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                >
                  Create Account
                </button>
              ) : (
                <button
                  type="submit"
                  className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-400 hover:bg-green-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                >
                  Log in
                </button>
              )}
            </div>

            {!signUp ? (
              <div className="text-center">
                <p className="text-sm text-gray-600">
                  Don't have an account?{" "}
                  <a
                    href="#"
                    onClick={isSignIn}
                    className="font-medium text-green-600 hover:text-green-500"
                  >
                    Sign up
                  </a>
                </p>
              </div>
            ) : (
              <div className="text-center">
                <p className="text-sm text-gray-600">
                  Already have an account?{" "}
                  <a
                    href="#"
                    onClick={() => setSignUp(() => false)}
                    className="font-medium text-green-600 hover:text-green-500"
                  >
                    Log In
                  </a>
                </p>
              </div>
            )}
          </form>
          <p className="mt-2 text-gray-600">{message}</p>
        </div>
      </div>
    </div>
  );
};

export default Login;
