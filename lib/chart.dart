import 'package:flutter/material.dart';

class Chart extends StatelessWidget {
  const Chart({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        actions: [
          Container(
             padding: const EdgeInsets.only(right: 116), 
            child: Row(
              children: [
                Image.asset('assets/images/LOGO-SU.png', width: 45, height: 45,),
                const SizedBox(width: 10,),
                const Text("Student Union", style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),), 
              ],
            )
          
      
          )
        ],

      ),
    );
  }
}