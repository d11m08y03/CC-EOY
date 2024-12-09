import 'package:eoy_frontend/footer.dart';
import 'package:eoy_frontend/home.dart';
import 'package:eoy_frontend/splash_screen.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'EOY 2024',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      initialRoute: '/',
      onGenerateRoute: (RouteSettings settings) {
        // Handle dynamic routes
        if (settings.name == '/home') {
          final token = settings.arguments as String; 
          return MaterialPageRoute(
            builder: (context) => Home(token: token),  
          );
        }

        if (settings.name == '/footer') {
          final token = settings.arguments as String; 
          return MaterialPageRoute(
            builder: (context) => Footer(token: token),  
          );
        }

        if (settings.name == '/') {
          return MaterialPageRoute(builder: (context) => const SplashScreen()); 
        }

        return null;
      },
    );
  }
}
