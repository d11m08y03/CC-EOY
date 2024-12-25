import 'package:eoy_frontend/environment.dart';
import 'package:flutter/material.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:barcode_scan2/barcode_scan2.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class Home extends StatefulWidget {
  final String token;
  const Home({super.key, required this.token});

  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  final TextEditingController _studentNumberController =
      TextEditingController();
  final _fullNameController = TextEditingController();
  final _emailController = TextEditingController();
  bool _isLoading = false;
  String _responseMessage = '';

  final String _apiUrl = '${Environment.serverUrl}:${Environment.port}';

  @override
  void dispose() {
    _studentNumberController.dispose();
    _fullNameController.dispose();
    _emailController.dispose();
    super.dispose();
  }

  Future<void> _submitStudentInfo() async {
    if (_studentNumberController.text.isEmpty) {
      setState(() {
        _responseMessage = "Student number is required.";
        print(widget.token);
      });
      return;
    }
    if (widget.token.isEmpty) {
      setState(() {
        _responseMessage = "Token is missing or invalid. Please log in again.";
      });
      return;
    } else {
      print("Token: ${widget.token}");
    }

    setState(() {
      _isLoading = true;
      _responseMessage = '';
    });
    var headers = {
      'Authorization': widget.token,
      'Content-Type': 'application/json',
    };
    var body = json.encode({
      "student_id": _studentNumberController.text,
    });
    try {
      var response = await http.put(
        Uri.parse("$_apiUrl/auth/students"),
        headers: headers,
        body: body,
      );

      var responseData = json.decode(response.body);

      if (response.statusCode == 200) {
        setState(() {
          _responseMessage = "Request successful";
        });

        String fullName = responseData['full_name'] ?? 'No Name Provided';
        String studentID = responseData['student_id'] ?? 'No ID Provided';
        String programOfStudy =
            responseData['program_of_study'] ?? 'No program_of_study Provided';
        String level = responseData['level'] ?? 'No level Provided';

        _showBottomModal(fullName, studentID, programOfStudy, level);
      } else if (response.statusCode == 404) {
        String errorMessage = responseData["error"] ?? "Student not found.";
        _show404AlertModal(errorMessage, _studentNumberController.text);
      } else {
        setState(() {
          var error = responseData["error"] ?? "Unknown error";
          _responseMessage = "Request failed: $error";
        });
      }
    } catch (e) {
      setState(() {
        _responseMessage = "An error occurred: $e";
      });
      print("Error details: $e");
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  void _show404AlertModal(String message, String studentID) {
    showDialog(
      context: context,
      barrierDismissible: false,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Row(
            children: [
              Icon(Icons.warning, color: Colors.orange),
              SizedBox(width: 10),
              Text("Warning"),
            ],
          ),
          content: Text(
            "$message.\n\nStudent not registered or is no more an student. Please Verify the student ID manually then deny or grant access.",
          ),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: const Text("No"),
            ),
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
                _showManualEntryModal(studentID);
              },
              child: const Text(
                "Yes",
                style: TextStyle(color: Colors.blue),
              ),
            ),
          ],
        );
      },
    );
  }

  void _showManualEntryModal(String studentID) {
    showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      builder: (BuildContext context) {
        return Padding(
          padding: EdgeInsets.only(
            bottom: MediaQuery.of(context).viewInsets.bottom,
            left: 20,
            right: 20,
            top: 20,
          ),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Row(
                children: [
                  Icon(Icons.person_add, color: Colors.blue),
                  SizedBox(width: 10),
                  Text(
                    "Manual Student Entry",
                    style: TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ],
              ),
              const Divider(height: 30, thickness: 2),
              Text(
                "Student ID: $studentID",
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w600,
                ),
              ),
              const SizedBox(height: 15),
              TextField(
                controller: _fullNameController,
                decoration: const InputDecoration(
                  labelText: 'Full Name *',
                  border: OutlineInputBorder(),
                ),
              ),
              const SizedBox(height: 15),
              TextField(
                controller: _emailController,
                decoration: const InputDecoration(
                  labelText: 'Email (Optional)',
                  border: OutlineInputBorder(),
                ),
                keyboardType: TextInputType.emailAddress,
              ),
              const SizedBox(height: 25),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  ElevatedButton(
                    onPressed: () {
                      Navigator.of(context).pop();
                    },
                    style: ElevatedButton.styleFrom(
                        backgroundColor:
                            const Color.fromARGB(255, 35, 119, 210),
                        foregroundColor:
                            const Color.fromARGB(255, 255, 255, 255)),
                    child: const Text("Back"),
                  ),
                  ElevatedButton(
                    onPressed: () {
                      _submitManualEntry(studentID);
                    },
                    style: ElevatedButton.styleFrom(
                        backgroundColor:
                            const Color.fromARGB(255, 35, 119, 210),
                        foregroundColor:
                            const Color.fromARGB(255, 255, 255, 255)),
                    child: const Text("Submit"),
                  ),
                ],
              ),
            ],
          ),
        );
      },
    );
  }

  void _showBottomModal(
      String fullName, String studentID, String programOfStudy, String level) {
    showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      backgroundColor: Colors.transparent,
      builder: (BuildContext context) {
        return Container(
          width: double.infinity,
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 20),
          decoration: const BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.vertical(
              top: Radius.circular(25),
            ),
            boxShadow: [
              BoxShadow(
                color: Colors.black26,
                blurRadius: 10.0,
                offset: Offset(0, -2),
              ),
            ],
          ),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Header
              const Row(
                children: [
                  Icon(
                    Icons.info_outline,
                    color: Colors.blue,
                    size: 28,
                  ),
                  SizedBox(width: 10),
                  Text(
                    "Student Information",
                    style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ],
              ),
              const Divider(
                height: 30,
                thickness: 2,
              ),

              _buildInfoRow(Icons.person, "Full Name", fullName),
              const SizedBox(height: 15),
              _buildInfoRow(Icons.badge, "Student ID", studentID),
              const SizedBox(height: 25),
              _buildInfoRow(Icons.school, "Program of Study", programOfStudy),
              const SizedBox(height: 15),
              _buildInfoRow(Icons.badge, "Level", level),
              const SizedBox(height: 25),
              // Close Button
              Center(
                child: ElevatedButton.icon(
                  onPressed: () {
                    Navigator.pop(context);
                  },
                  icon: const Icon(Icons.close),
                  label: const Text("Close"),
                  style: ElevatedButton.styleFrom(
                    foregroundColor: Colors.white,
                    backgroundColor: Colors.redAccent,
                    padding: const EdgeInsets.symmetric(
                        horizontal: 20, vertical: 12),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                ),
              ),
            ],
          ),
        );
      },
    );
  }

  Widget _buildInfoRow(IconData icon, String title, String value) {
    return Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Icon(
          icon,
          color: Colors.blueAccent,
          size: 24,
        ),
        const SizedBox(width: 15),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                title,
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w600,
                  color: Colors.black54,
                ),
              ),
              const SizedBox(height: 5),
              Text(
                value,
                style: const TextStyle(
                  fontSize: 16,
                  color: Colors.black87,
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }

  Future<void> _checkAndRequestPermission() async {
    var status = await Permission.camera.status;
    if (!status.isGranted) {
      await Permission.camera.request();
    }
  }

  Future<void> _scanBarcode() async {
    await _checkAndRequestPermission();

    try {
      final result = await BarcodeScanner.scan();

      if (result.type == ResultType.Barcode) {
        setState(() {
          _studentNumberController.text = result.rawContent;
        });

        await _submitStudentInfo();
      } else {
        setState(() {
          _responseMessage = "Scan failed or was canceled";
        });
      }
    } catch (e) {
      setState(() {
        _responseMessage = "Error occurred while scanning: $e";
      });
    }
  }

  void _showStudentData(BuildContext context) {
    showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      shape: const RoundedRectangleBorder(
        borderRadius: BorderRadius.vertical(top: Radius.circular(20)),
      ),
      builder: (context) {
        return DraggableScrollableSheet(
          expand: false,
          initialChildSize: 0.6,
          minChildSize: 0.4,
          maxChildSize: 0.8,
          builder: (context, scrollController) {
            return Container(
              padding: const EdgeInsets.all(16),
              decoration: const BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.vertical(top: Radius.circular(20)),
              ),
              child: ListView(
                controller: scrollController,
                children: [
                  const Text(
                    'Sliding Screen',
                    style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
                  ),
                  const SizedBox(height: 16),
                  const Text('This is the content of the sliding screen.'),
                  const SizedBox(height: 16),
                  ElevatedButton(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    child: const Text('Close'),
                  ),
                ],
              ),
            );
          },
        );
      },
    );
  }

  void _submitManualEntry(String studentID) async {
    String fullName = _fullNameController.text.trim();
    String email = _emailController.text.trim();

    if (fullName.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text("Full Name is required."),
          backgroundColor: Colors.red,
        ),
      );
      return;
    }

    Map<String, dynamic> studentData = {
      "student_id": studentID,
      "full_name": fullName,
      "email": email.isNotEmpty ? email : null,
    };

    String jsonString = json.encode(studentData);

    setState(() {
      _isLoading = true; // Show loading indicator
    });

    var headers = {
      'Authorization': widget.token,
      'Content-Type': 'application/json',
    };

    final String postUrl = "$_apiUrl/auth/students";

    try {
      var response = await http.post(
        Uri.parse(postUrl),
        headers: headers,
        body: jsonString,
      );

      var responseData = json.decode(response.body);

      if (response.statusCode == 200) {
        showDialog(
          context: context,
          builder: (BuildContext context) {
            return AlertDialog(
              title: const Text("Success"),
              content: const Text(
                  "Student information has been submitted successfully."),
              actions: [
                TextButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                    Navigator.of(context).pop();
                  },
                  child: const Text("OK"),
                ),
              ],
            );
          },
        );
      } else {
        String errorMessage =
            responseData["error"] ?? "Unknown error occurred.";
        showDialog(
          context: context,
          builder: (BuildContext context) {
            return AlertDialog(
              title: const Text("Error"),
              content: Text("Submission failed: $errorMessage"),
              actions: [
                TextButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                    Navigator.of(context).pop();
                  },
                  child: const Text("OK"),
                ),
              ],
            );
          },
        );
      }
    } catch (e) {
      showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: const Text("Error"),
            content: Text("An error occurred: $e"),
            actions: [
              TextButton(
                onPressed: () {
                  Navigator.of(context).pop();
                  Navigator.of(context).pop();
                },
                child: const Text("OK"),
              ),
            ],
          );
        },
      );
    } finally {
      setState(() {
        _isLoading = false;
      });

      // Clear the input fields
      _fullNameController.clear();
      _emailController.clear();
    }
  }

  @override
  Widget build(BuildContext context) {
    // Get screen size for responsiveness
    final screenWidth = MediaQuery.of(context).size.width;
    final screenHeight = MediaQuery.of(context).size.height;

    return Scaffold(
      appBar: AppBar(
        actions: [
          Container(
            padding: EdgeInsets.only(right: screenWidth * 0.3),
            child: Row(
              children: [
                Image.asset('assets/images/LOGO-SU.png',
                    width: screenWidth * 0.1, height: screenWidth * 0.1),
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
          )
        ],
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: screenWidth * 0.05),
          child: Column(
            children: [
              SizedBox(
                height: 40,
              ),
              Center(
                child: ClipRRect(
                  borderRadius: BorderRadius.circular(20),
                  child: Image.asset(
                    'assets/images/SID.png',
                    width: screenWidth * 0.9,
                    fit: BoxFit.cover,
                  ),
                ),
              ),
              SizedBox(height: screenHeight * 0.05),
              SizedBox(
                width: screenWidth * 0.9,
                child: TextField(
                  controller: _studentNumberController,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    labelText: 'Student ID Number',
                  ),
                ),
              ),
              SizedBox(height: screenHeight * 0.1),
              SizedBox(
                width: screenWidth * 0.8,
                height: 55,
                child: ElevatedButton(
                  onPressed: _isLoading ? null : _submitStudentInfo,
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.black,
                    foregroundColor: Colors.white,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                  child: const Text(
                    'Mark as present',
                    style: TextStyle(
                      fontSize: 15,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ),
              ),
              SizedBox(height: screenHeight * 0.05),
              SizedBox(
                width: screenWidth * 0.8,
                height: 55,
                child: ElevatedButton(
                  onPressed: () {
                    _scanBarcode();
                  },
                  style: ElevatedButton.styleFrom(
                    backgroundColor: const Color.fromARGB(255, 35, 119, 210),
                    foregroundColor: Colors.white,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                  child: const Text(
                    'Scan Barcode',
                    style: TextStyle(
                      fontSize: 15,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ),
              ),
              SizedBox(height: screenHeight * 0.03),
              Text(
                _responseMessage,
                style: TextStyle(
                  color: _responseMessage.contains("successful")
                      ? Colors.green
                      : (_responseMessage.contains("already registered")
                          ? Colors.orange
                          : Colors.red),
                ),
                textAlign: TextAlign.center,
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget logoWidget(String imageName) {
    return ClipRRect(
      borderRadius: BorderRadius.circular(20),
      child: Image.asset(
        imageName,
        fit: BoxFit.fitWidth,
      ),
    );
  }
}
