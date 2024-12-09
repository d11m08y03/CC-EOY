import 'dart:convert';
import 'package:eoy_frontend/environment.dart';
import 'package:eoy_frontend/footer.dart';
import 'package:eoy_frontend/home.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

bool _isAdmin = false;

class Login extends StatefulWidget {
  const Login({super.key});

  @override
  State<Login> createState() => _LoginState();
}

class _LoginState extends State<Login> {
  final _formKey = GlobalKey<FormState>();
  bool _isLoading = false;

  final String _apiUrl = '${Environment.serverUrl}:${Environment.port}';

  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();

  Future login() async {
    if (!_formKey.currentState!.validate()) {
      return;
    }

    setState(() {
      _isLoading = true;
    });

    final email = emailController.text.trim().toLowerCase();
    final password = passwordController.text;

    try {
      var headers = {'Content-Type': 'application/json'};
      var data = json.encode({"email": email, "password": password});

      var url = Uri.parse("$_apiUrl/login");
      var response = await http.post(url, headers: headers, body: data);

      if (response.statusCode == 200) {
        final responseData = json.decode(response.body);
        final token = responseData['token'];
        _isAdmin = responseData['is_admin'];

        if (token != null) {
          if (_isAdmin) {
            Navigator.push(
              context,
              MaterialPageRoute(builder: (context) => Footer(token: token)),
            );
          } else {
            Navigator.push(
              context,
              MaterialPageRoute(builder: (context) => Home(token: token)),
            );
          }
        } else {
          _showErrorDialog("Login failed: No token received from server.");
        }
      } else if (response.statusCode == 401) {
        final responseData = json.decode(response.body);
        final errorMessage = responseData['message'] ??
            ' Unauthorized user. Please contact Zakariya.';
        _showErrorDialog("Login failed: $errorMessage");
      } else {
        _showErrorDialog("Login failed: ${response.statusCode}");
      }
    } catch (e) {
      _showErrorDialog("An error occurred: $e, Url: $_apiUrl");
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  void _showErrorDialog(String message) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text("Error", style: TextStyle(color: Colors.black)),
        content: Text(
          message,
          style: const TextStyle(color: Colors.red, fontSize: 16),
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.of(context).pop(),
            child: const Text("OK"),
          ),
        ],
      ),
    );
  }

  String? _validateEmail(String? value) {
    if (value == null || value.isEmpty) {
      return "Email is required";
    }
    final emailRegex = RegExp(r'^[^@]+@[^@]+\.[^@]+$');
    if (!emailRegex.hasMatch(value)) {
      return "Enter a valid email address";
    }
    return null;
  }

  String? _validatePassword(String? value) {
    if (value == null || value.isEmpty) {
      return "Password is required";
    }
    if (value.length < 8) {
      return "Password must be at least 8 characters long";
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.symmetric(vertical: 50, horizontal: 20),
            child: Form(
              key: _formKey,
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Image.asset(
                    'assets/images/LOGO-SU.png',
                    height: MediaQuery.of(context).size.width * 0.3,
                    width: MediaQuery.of(context).size.width * 0.3,
                  ),
                  const SizedBox(height: 30),
                  Text(
                    "Please log in",
                    style: TextStyle(
                      fontSize: MediaQuery.of(context).size.width * 0.08,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  const SizedBox(height: 20),
                  TextFormField(
                    controller: emailController,
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: 'Email',
                    ),
                    validator: _validateEmail,
                  ),
                  const SizedBox(height: 20),
                  TextFormField(
                    controller: passwordController,
                    obscureText: true,
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: 'Password',
                    ),
                    validator: _validatePassword,
                  ),
                  const SizedBox(height: 40),
                  SizedBox(
                    width: double.infinity,
                    height: 55,
                    child: ElevatedButton(
                      onPressed: _isLoading ? null : login,
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.black,
                        foregroundColor: Colors.white,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(10),
                        ),
                      ),
                      child: _isLoading
                          ? const CircularProgressIndicator(color: Colors.white)
                          : const Text('Log in'),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
