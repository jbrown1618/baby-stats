import { ThemedText } from "@/components/ThemedText";
import { Link, Tabs, useLocalSearchParams } from "expo-router";

export default function TabsLayout() {
  const { babyID } = useLocalSearchParams();
  return (
    <>
      <Tabs>
        <Tabs.Screen
          name="index"
          options={{ headerShown: false, href: `/${babyID}` }}
        />
        <Tabs.Screen
          name="food"
          options={{ headerShown: false, href: `/${babyID}/food` }}
        />
        <Tabs.Screen
          name="sleep"
          options={{ headerShown: false, href: `/${babyID}/sleep` }}
        />
      </Tabs>
      <Link
        href={`/${babyID}/add-event`}
        style={{
          position: "absolute",
          right: 10,
          bottom: 60,
          backgroundColor: "red",
          padding: 10,
        }}
      >
        <ThemedText style={{ color: "white" }}>Add Event</ThemedText>
      </Link>
    </>
  );
}
