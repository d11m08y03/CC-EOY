import 'dart:convert';

import 'package:eoy_frontend/environment.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class AddOrganizer extends StatefulWidget {
  final String token;
  const AddOrganizer({super.key, required this.token});

  @override
  State<AddOrganizer> createState() => _AddOrganizerState();
}

class _AddOrganizerState extends State<AddOrganizer> {
  TextEditingController nameController = TextEditingController();
  TextEditingController organizerEmailController = TextEditingController();
  TextEditingController organizerPasswordController = TextEditingController();

  void add(String name, String email, String password) async {
    print(widget.token);
    try {
      final response = await http.post(
        Uri.parse('${Environment.serverUrl}:${Environment.port}/register'),
        headers: {
          'Authorization': widget.token,
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          'name': name,
          'email': email,
          'password': password,
        }),
      );

      if (response.statusCode == 201) {
        print("Account created successfully");
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text("Account created successfully")),
        );
      } else {
        print("Failed to create account. Status Code: ${response.statusCode}");
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text("Failed to create account")),
        );
      }
    } catch (e) {
      print("Error: ${e.toString()}");
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text("Error: ${e.toString()}")),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        actions: [
          Container(
            padding: const EdgeInsets.only(right: 116),
            child: Row(
              children: [
                Image.asset(
                  'assets/images/LOGO-SU.png',
                  width: 45,
                  height: 45,
                ),
                const SizedBox(width: 10),
                const Text(
                  "Student Union",
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
      body: SingleChildScrollView(
        child: Column(
          children: [
            const SizedBox(height: 140),
            const Center(
              child: Text(
                "Add new organizer",
                style: TextStyle(
                  fontSize: 30,
                  fontWeight: FontWeight.w800,
                ),
              ),
            ),
            const SizedBox(height: 30),
            SizedBox(
              width: 350,
              child: TextField(
                controller: nameController,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Name',
                ),
              ),
            ),
            const SizedBox(height: 10),
            SizedBox(
              width: 350,
              child: TextField(
                controller: organizerEmailController,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Email',
                ),
              ),
            ),
            const SizedBox(height: 10),
            SizedBox(
              width: 350,
              child: TextField(
                controller: organizerPasswordController,
                decoration: const InputDecoration(
                  border: OutlineInputBorder(),
                  labelText: 'Password',
                ),
                obscureText: true,
              ),
            ),
            const SizedBox(height: 40),
            SizedBox(
              width: 300,
              height: 55,
              child: ElevatedButton(
                onPressed: () {
                  add(
                    nameController.text.trim(),
                    organizerEmailController.text.trim(),
                    organizerPasswordController.text.trim(),
                  );
                },
                style: ElevatedButton.styleFrom(
                  backgroundColor: const Color.fromARGB(255, 35, 119, 210),
                  foregroundColor: Colors.white,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                ),
                child: const Text(
                  'Add organizer',
                  style: TextStyle(
                    fontSize: 15,
                    fontWeight: FontWeight.w700,
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
