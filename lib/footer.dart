import 'package:eoy_frontend/add_organizer.dart';
import 'package:eoy_frontend/chart.dart';
import 'package:eoy_frontend/home.dart';
import 'package:flutter/material.dart';

class Footer extends StatefulWidget {
  final String token; 

  const Footer({super.key, required this.token}); 

  @override
  State<Footer> createState() => _FooterState();
}

class _FooterState extends State<Footer> {
  int currentIndex = 1;
  

  @override
  Widget build(BuildContext context) {
     final pages = [
      AddOrganizer(token: widget.token),
      Home(token: widget.token), 
      const Chart(),
    ];
    return Scaffold(
      
      body: pages[currentIndex],
    
      bottomNavigationBar: BottomNavigationBar(
        items: [
          BottomNavigationBarItem(
            
            icon: Icon(Icons.add_circle_rounded,
                color: currentIndex == 0
                    ? const Color.fromARGB(255, 35, 119, 210)
                    : Colors.grey),
            label: 'Add organizer',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.credit_card_rounded,
                color: currentIndex == 1
                    ? const Color.fromARGB(255, 35, 119, 210)
                    : Colors.grey),
            label: 'Scan',
          ),
          // BottomNavigationBarItem(
          //   icon: Icon(Icons.data_exploration,
          //       color: currentIndex == 2
          //           ? const Color.fromARGB(255, 35, 119, 210)
          //           : Colors.grey),
          //   label: 'Data Analysis',
          // ),
        ],
        currentIndex: currentIndex,
        onTap: (index) {
          setState(() {
            currentIndex = index;
          });
        },
      ),
    );
    
  }

 

}
